<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method ="post">
      用户名:<input type="text" name="username">
      密码:<input type="password" name="password">
      <input type="hidden" name="token" value="{{.}}"><!--防止表单重复提交-->
      <input type="submit" value="登陆">
</form>
</body>
</html>