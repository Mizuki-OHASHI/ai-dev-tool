<?php

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;

class HelloController extends Controller
{
    public function world()
    {
        $data = ['name' => 'World'];
        return view('hello', $data);
    }

    public function data($name)
    {
        $data = ['name' => $name];
        return view('hello', $data);
    }
}
