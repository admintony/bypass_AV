package main

import (
  "fmt"
  "syscall"
  "unsafe"
)

// 定义Windows API函数的参数类型和返回值类型
// 下面是NetLocalGroupAddMembers函数的定义：
// NET_API_STATUS NetLocalGroupAddMembers(
//     LPCWSTR servername,
//     LPCWSTR groupname,
//     DWORD   level,
//     LPBYTE  buf,
//     DWORD   totalentries
// );
var (
    mod         = syscall.NewLazyDLL("netapi32.dll")
    proc        = mod.NewProc("NetLocalGroupAddMembers")
    group       = syscall.StringToUTF16Ptr("administrators") // 组名
    server      = syscall.StringToUTF16Ptr("")               // 服务器名，为空表示本地计算机
    level       = uint32(3)                                  // 参数level的值
    totalEntries = uint32(1)                                  // 参数totalentries的值
)

// 定义LOCALGROUP_MEMBERS_INFO_3
type localGroupMembersInfo3 struct {
    Lgrmi3_domainandname *uint16 // 用户名
}

var (
  modNetapi32 = syscall.NewLazyDLL("netapi32.dll")
  procNetUserAdd = modNetapi32.NewProc("NetUserAdd")
  procNetLocalGroupAddMembers = modNetapi32.NewProc("NetLocalGroupAddMembers")
)

// USER_INFO_1 contains information about a user account.
type USER_INFO_1 struct {
  Name            *uint16
  Password        *uint16
  PasswordAge     uint32
  Priv            uint32
  HomeDir         *uint16
  Comment         *uint16
  Flags           uint32
  ScriptPath      *uint16
}

func main() {
  // Set up the USER_INFO_1 struct with the user's information.
  user := USER_INFO_1{
    Name: syscall.StringToUTF16Ptr("newuser"),
    Password: syscall.StringToUTF16Ptr("password"),
    PasswordAge: 0,
    Priv: 1,
    HomeDir: nil,
    Comment: nil,
    Flags: 0x0001,
    ScriptPath: syscall.StringToUTF16Ptr("C:\\Users\\newuser\\profile.cmd"),
  }

  // Call the NetUserAdd function to add the user to the system.
  var parmErr uint32
  var err error
  r1, _, err := procNetUserAdd.Call(
    0,
    1,
    uintptr(unsafe.Pointer(&user)),
    uintptr(unsafe.Pointer(&parmErr)),
  )
  if r1 != 0 {
    // An error occurred.
    fmt.Println("Error adding user:", err)
  } else {
    fmt.Println("Successfully added user.")

    // 将newuser添加到管理员组
    username := syscall.StringToUTF16Ptr("newuser") // 用户名
    buf := localGroupMembersInfo3{Lgrmi3_domainandname: username}
    r1, _, err := proc.Call(
        uintptr(unsafe.Pointer(server)),
        uintptr(unsafe.Pointer(group)),
        uintptr(level),
        uintptr(unsafe.Pointer(&buf)),
        uintptr(totalEntries),
    )
    if r1 != 0 {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Successfully added administrators group.")
    }
  }
}
