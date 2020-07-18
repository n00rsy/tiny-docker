package container

import (
   "fmt"
   "os"
   "os/exec"
   "syscall"
)


func Init(){

   // cmd to call ubernetes with new args
   // calling /proc/self/exe calls ubernetes.go. same as fork/exec
   cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   cmd.Stdin = os.Stdin

   // setting flags for new process
   cmd.SysProcAttr = &syscall.SysProcAttr{
      Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
      Unshareflags: syscall.CLONE_NEWNS,
   }

   must(cmd.Run())

}

func Run(){
   fmt.Printf("Running %v \n", os.Args[2:])

   config_cgroup()
   cmd := exec.Command(os.Args[2], os.Args[3:]...)
   cmd.Stdin = os.Stdin
   cmd.Stdout = os.Stdout
   cmd.Stderr = os.Stderr
   
   //must(os.MkdirAll("rootfs/oldrootfs", 0700))
   must(syscall.Sethostname([]byte("container")))
   must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
   must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
   must(os.Chdir("/"))
   // run that bish
   if err := cmd.Run(); err != nil {
      fmt.Println("ERROR", err)
      os.Exit(1)
   }
}

func config_cgroup(){


}

func must(err error){
   if err != nil {
      panic(err)
   }
}
