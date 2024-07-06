package main

deny[msg] { # ワークフローレベルのPermissionが省略されていたら拒否
  not input.permissions
  msg = "Workflow permission are missing"
}

deny [msg]{ # ワークフローレベルのpermissionsで空以外が指定されていたら拒否
  input.permissions != {}
  msg = sprintf("Workflow permissions are not empty: %v",[input.permissions])
}
