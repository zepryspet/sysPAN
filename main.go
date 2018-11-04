package main

import(
  "github.com/RackSec/srslog"
  "log"
  "bufio"
  "os"
  "crypto/tls"
  "flag"
  //"fmt"
)

func main (){
  tr := flag.String("tr", "udp", "either udp or ssl")
  ip := flag.String("ip", "192.168.1.1", "user id agent IP address")
  flag.Parse()
  //fmt.Println("tr:", *tr)
  //fmt.Println("ip:", *ip)
  fname := "logs.txt"
  if *tr == "udp"{
    sysUDP(*ip, fname)
  }else if *tr == "ssl"{
    sysTLS(*ip, fname)
  }else{
    log.Fatal("Invalid transport layer, please use either udp or ssl")
  }
}

func sysTLS(ip string, fname string){
  config := tls.Config{
    InsecureSkipVerify: true,
  }
  w, err := srslog.DialWithTLSConfig("tcp+tls", ip+":6514", srslog.LOG_ERR, "testtag", &config)
  if err != nil {
      log.Fatal("failed to connect to syslog:", err)
  }
  defer w.Close()
  SendData(fname,  w)
}

func sysUDP(ip string, fname string){
  w, err := srslog.Dial("udp", ip+":514", srslog.LOG_ERR, "testtag")
  if err != nil {
      log.Fatal("failed to connect to syslog:", err)
  }
  defer w.Close()
  SendData(fname,  w)
}

func SendData(fname string, w *srslog.Writer){
  file, err := os.Open(fname)
  if err != nil {
     log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
     w.Info(scanner.Text())
  }
}
