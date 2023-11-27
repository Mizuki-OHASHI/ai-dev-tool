<?php

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class ValidationController extends Controller
{
    public function showInput()
    {
        return view('playground.validation');
    }
    public function addData(Request $request)
    {
        $request->validate([
            'name'  => 'required',
            'age'   => 'required|numeric',
            'email' => 'required|email',
        ]);
        return '<h1>Input Done</h1>';
    }
}
