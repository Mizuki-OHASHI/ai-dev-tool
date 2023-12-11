<!DOCTYPE html>
<html lang="ja">
<head>
	<meta charset="UTF-8">
	<title>Hello World by Blade</title>
  <style>
    .blue-500 {
      color: #2196F3;
    }
  </style>
</head>
	<body>
		<h1 class="blue-500">Hello {{$name}}! Blade!</h1>
    <h2>現在の日時: {{date("Y/m/d H:i:s")}}</h2>
	</body>
</html>