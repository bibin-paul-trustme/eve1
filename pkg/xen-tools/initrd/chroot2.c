#define _GNU_SOURCE
#include <sched.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <sys/mount.h>
#include <signal.h>

typedef struct clone_args clone_args;

struct clone_args {
    char **args;
    char *command;
    uid_t uid, gid;
};

#define STACK_SIZE (8 * 1024 * 1024)
static char child_stack[STACK_SIZE];    /* Space for child's stack */

static int childFunc(void *args)
{
    clone_args *parsed_args = (clone_args *)args;
    mount("proc", "/proc", "proc", 0, NULL);

    setgid(parsed_args->gid);
    setuid(parsed_args->uid);

    execvp(parsed_args->command, parsed_args->args);
}

int main(int argc, char **argv) {
    uid_t uid, gid;
    char *endptr;
    pid_t child_pid;
    struct clone_args args;

    setsid();
    ioctl(0, TIOCSCTTY, 1);

    chroot(argv[1]);
    chdir(argv[2]);

    uid = strtol(argv[3], &endptr, 10);
    gid = strtol(argv[4], &endptr, 10);

    args.uid = uid;
    args.gid = gid;
    args.command = argv[5];
    args.args = argv + 5;

    child_pid = clone(childFunc, child_stack + STACK_SIZE, CLONE_NEWPID | SIGCHLD, (void *)(&args));

    waitpid(child_pid, NULL, 0);
}
