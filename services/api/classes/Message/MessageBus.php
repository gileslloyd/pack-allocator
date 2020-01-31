<?php

declare(strict_types=1);

namespace Message;

use PhpAmqpLib\Connection\AMQPStreamConnection;
use PhpAmqpLib\Message\AMQPMessage;

class MessageBus
{
    /**
     * @var string
     */
    private $host;

    /**
     * @var AMQPStreamConnection
     */
    private $connection;

    /**
     * @var \PhpAmqpLib\Channel\AMQPChannel
     */
    private $channel;

    /**
     * @var string|null
     */
    private $correlation_id = null;

    /**
     * @var array|null
     */
    private $response = null;

    public function __construct(string $host, int $port, string $user, string $password)
    {
        $this->host = $host;

        $this->connection = new AMQPStreamConnection(
            $host,
            $port,
            $user,
            $password
        );

        $this->channel = $this->connection->channel();
    }

    public function async(string $queue, array $message): void
    {
        $this->channel->queue_declare($queue, false, false, false, false);

        $this->channel->basic_publish(new AMQPMessage(json_encode($message)), '', $queue);
    }

    public function sync(string $queue, array $message): MessagePayload
    {
        $this->response = null;
        $this->correlation_id = uniqid();

        list($callback_queue, ,) = $this->channel->queue_declare("", false, false, true, false);
        $this->channel->basic_consume($callback_queue, '', false, false, false, false, [$this, 'onResponse']);

        $this->channel->basic_publish(
            new AMQPMessage(
                json_encode($message),
                ['correlation_id' => $this->correlation_id, 'reply_to' => $callback_queue]
            ),
            '',
            $queue
        );

        while(!$this->response) {
            $this->channel->wait();
        }

        return new MessagePayload($this->response);
    }

    public function publish(string $exchange, array $message)
    {
        $this->channel->exchange_declare(
            $exchange,
            'fanout',
            false,
            false,
            false
        );

        $this->channel->basic_publish(new AMQPMessage(json_encode($message)), $exchange);
    }

    public function onResponse(AMQPMessage $response) {
        if($response->get('correlation_id') == $this->correlation_id) {
            $response = json_decode($response->body, true);
            $this->response = $response['payload'] ?? $response;
        }
    }

    public function __destruct()
    {
        $this->connection->close();
    }
}
