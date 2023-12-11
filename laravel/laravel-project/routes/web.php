<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\HelloController;
use App\Http\Controllers\SandController;
use App\Http\Controllers\ValidationController;

/*
|--------------------------------------------------------------------------
| Web Routes
|--------------------------------------------------------------------------
|
| Here is where you can register web routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "web" middleware group. Make something great!
|
*/

Route::get('/', function () {
    return view('welcome');
});
Route::get('hello', function () {
    print("<h1>Hello World!</h1>");
    return null;
});
Route::get('hello/{name}', function ($name) {
    print("<h1>Hello $name!</h1>");
    return null;
});
Route::redirect('uttc', 'https://www.uttc.dev');
Route::get('helloBlade', function () {
    $data = ['name' => 'World'];
    return view('hello', $data);
});
Route::get('if', function () {
    $data = ['random' => rand(0,100)];
    return view('playground.if', $data);
});
Route::get('loop', function () {
    $data = ['array' => ['a', 'b', 'c', 'd', 'e'], 'random' => rand(0,100)];
    return view('playground.loop', $data);
});

Route::get('helloController', [HelloController::class, 'world']);
Route::get('helloController/{name}', [HelloController::class, 'data']);
Route::get('middlewareTest', function () {
    return '<p>Middleware Test: This is request process.</p>';
})->middleware("record-ip-address:123");

Route::get('/sand/newBook', [SandController::class, 'newBook']);
Route::get('/sand/newBook2', [SandController::class, 'newBook2']);
Route::get('/sand/newBook3', [SandController::class, 'newBook3']);
Route::get('/sand/newMathNote', [SandController::class, 'newMathNote']);

Route::get('playground/validation', [ValidationController::class, 'showInput']);
Route::post('playground/addData', [ValidationController::class, 'addData']);

Route::get('playground/sql', [SandController::class, 'runningRawSQLQueries']);