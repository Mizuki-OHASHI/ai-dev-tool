<!DOCTYPE html>
<html lang="ja">
    <head>
        <meta charset="UTF-8">
        <title>Validation</title>
        <link href="{{ asset('css/app.css') }}" rel="stylesheet">
        <style>
            form {
                width: 500px;
                margin: 20px auto;
            }

            .form-label {
                display: block;
                margin-bottom: 5px;
                font-weight: bold;
            }
            
            .form-input {
                border: 1px solid #ccc;
                border-radius: 4px;
                padding: 8px;
                margin-bottom: 10px;
                width: 100%;
            }
            
            .form-textarea {
                border: 1px solid #ccc;
                border-radius: 4px;
                padding: 8px;
                margin-bottom: 10px;
                width: 100%;
            }
            
            .form-submit {
                background-color: #007bff;
                color: #fff;
                border: none;
                border-radius: 4px;
                padding: 10px 20px;
                cursor: pointer;
            }
            
            .form-submit:hover {
                background-color: #0056b3;
            }
        </style>
    </head>
    <body>
        <!--
        @if($errors->any())
            <section style="border: 1px red solid">
                <ul style="color: red;">
                    @foreach($errors->all() as $error)
                        <li>{{ $error }}</li>
                    @endforeach
                </ul>
            </section>
        @endif
        -->

        <form method="POST" action="/playground/addData" method="post">
            @csrf

            <label for="name" class="form-label">
                Name:
                <input type="text" id="name" name="name" class="form-input">
            </label>
            @error('name')
                <p style="color: red;">{{ $message }}</p>
            @enderror
            
            <label for="age" class="form-label">
                Age:
                <input type="text" id="age" name="age" class="form-input">
            </label>
            @error('age')
                <p style="color: red;">{{ $message }}</p>
            @enderror
            
            <label for="email" class="form-label">
                email:
                <input type="text" id="email" name="email" class="form-input">
            </label>
            @error('email')
                <p style="color: red;">{{ $message }}</p>
            @enderror
            
            <label for="comment" class="form-label">Comment:</label>
            <textarea id="comment" name="comment" rows="4" cols="50" class="form-textarea"></textarea>
            
            <button type="submit" class="form-submit">Submit</button>
        </form>
    </body>
</html>
