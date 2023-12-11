<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="UTF-8">
  <title>loop in template</title>
</head>
<body>
  <h1>loop in template</h1>
  <h2>for loop</h2>
  <ul>
    @for ($i = 0; $i < 5; $i++)
      <li>{{$i}}</li>
    @endfor
  </ul>
  <h2>foreach loop</h2>
  <ul>
    @foreach ($array as $item)
      @if ($loop->first)
        <li>first: {{$item}}</li>
      @elseif ($loop->last)
        <li>last: {{$item}}</li>
      @else
        <li>{{$loop->iteration}}: {{$item}}</li>
      @endif
    @endforeach
  </ul>
  <h2>forelse loop</h2>
  <ul>
    @forelse ($array as $item)
      <li>{{$item}}</li>
    @empty
      <li>empty</li>
    @endforelse
  </ul>
  <h2>while loop</h2>
  <ul>
    @while ($random < 90)
      <li>{{$random}}</li>
      <?php $random = mt_rand(0, 99); ?>
    @endwhile
  </ul>
</body>