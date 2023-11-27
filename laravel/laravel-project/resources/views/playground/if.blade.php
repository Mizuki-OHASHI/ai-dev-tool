<!DOCTYPE html>
<html lang="ja">
<head>
	<meta charset="UTF-8">
	<title>if/switch in template</title>
</head>
<body>
  <h1>if in template</h1>
  @if ($random < 50)
    <h2>{{$random}} is less than 50</h2>
  @elseif ($random < 70)
    <h2>{{$random}} is less than 70</h2>
  @else
    <h2>{{$random}} is greater than 70</h2>
  @endif
  <h1>switch in template</h1>
  @switch(round($random / 25))
    @case(0)
      <h2>{{round($random / 25)}} is 0</h2>
      @break
    @case(1)
      <h2>{{round($random / 25)}} is 1</h2>
      @break
    @case(2)
      <h2>{{round($random / 25)}} is 2</h2>
      @break
    @default
      <h2>{{round($random / 25)}} is not 0, 1, or 2</h2>
  @endswitch
</body>
