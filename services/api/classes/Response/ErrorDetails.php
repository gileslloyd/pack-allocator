<?php
declare(strict_types=1);

namespace Response;

class ErrorDetails implements ToArrayInterface
{

	use ToArrayTrait;

	/**
	 * @var string
	 */
	protected $error_code;

	/**
	 * @var string
	 */
	protected $message;

	/**
	 * @var string
	 */
	protected $details;

	/**
	 * @var string
	 */
	protected $more_info;

	public function __construct(string $error_code, string $message, string $details = '', string $more_info = '')
	{
		$this->error_code = $error_code;
		$this->message = $message;
		$this->details = $details;
		$this->more_info = $more_info;
	}

}
