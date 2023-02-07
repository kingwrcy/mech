package main

import (
   "2a.pages.dev/mech"
   "2a.pages.dev/mech/amc"
   "flag"
   "os"
   "path/filepath"
)

type flags struct {
   address string
   bandwidth int64
   email string
   mech.Stream
   password string
   verbose bool
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   var f flags
   // a
   flag.StringVar(&f.address, "a", "", "address")
   // c
   f.Client_ID = filepath.Join(home, "mech/client_id.bin")
   flag.StringVar(&f.Client_ID, "c", f.Client_ID, "client ID")
   // e
   flag.StringVar(&f.email, "e", "", "email")
   // f
   flag.Int64Var(&f.bandwidth, "f", 2_532_000, "video bandwidth")
   // i
   flag.BoolVar(&f.Info, "i", false, "information")
   // k
   f.Private_Key = filepath.Join(home, "mech/private_key.pem")
   flag.StringVar(&f.Private_Key, "k", f.Private_Key, "private key")
   // p
   flag.StringVar(&f.password, "p", "", "password")
   // v
   flag.BoolVar(&f.verbose, "v", false, "verbose")
   flag.Parse()
   if f.verbose {
      amc.Client.Log_Level = 2
   }
   if f.email != "" {
      err := f.login()
      if err != nil {
         panic(err)
      }
   } else if f.address != "" {
      err := f.download()
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
