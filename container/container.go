package container

import (
   "fmt"
   "os"
   "os/exec"
   "syscall"
   "path/filepath"
   "io/ioutil"
   "strconv"
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
   

   fmt.Printf("Running %v \n", os.Args[2:])
   // run that bish
   if err := cmd.Run(); err != nil {
      fmt.Println("ERROR", err)
      os.Exit(1)
   }
}

func config_cgroup(){
   cgroups := "/sys/fs/cgroup/"
   pids := filepath.Join(cgroups, "pids")
   os.Mkdir(pids, 0755)
   os.Mkdir(filepath.Join(pids, "container"), 0755)
   fmt.Println("created directory: ", filepath.Join(pids, "container"))

   // remove new cgroup after container exists - not sure exactly how this works
   must(ioutil.WriteFile(filepath.Join(pids, "container/notifiy_on_release"), []byte("1"), 0700))
   must(ioutil.WriteFile(filepath.Join(pids, "container/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))

}

func must(err error){
   if err != nil {
      panic(err)
   }
}
