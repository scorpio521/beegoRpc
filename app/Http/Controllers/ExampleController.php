<?php

namespace App\Http\Controllers;
use Illuminate\Http\Request;
use Greeter\GreeterClient;
use Document\DocumentClient;
use Document\DocumentRequest;

class ExampleController extends Controller
{
    /**
     * Create a new controller instance.
     *
     * @return void
     */
//    public function __construct()
//    {
//    }
    public function getRpc(Request $request)
    {

        $host = 'localhost:50051';
        $client = new DocumentClient($host,[
            'credentials' => (new \Grpc\ChannelCredentials())::createInsecure(),
        ]);


            $name = '30010';
        $request = new DocumentRequest();
        $request->setName($name);
        $call = $client->SayHello($request);
        list($response, $status) = $call->wait();
        var_dump('message = '.$response->getMessage());

    }
}
