class ApiError {
  status: number;

  message: string;

  constructor(status: number, message: string) {
    this.status = status;
    this.message = message;
  }

  static unauthorized() {
    return new ApiError(401, 'Unauthorized');
  }

  static forbidden() {
    return new ApiError(403, 'Forbidden');
  }

  static notFound() {
    return new ApiError(404, 'Not Found');
  }

  static internal() {
    return new ApiError(500, 'Internal Server Error');
  }
}

export default ApiError;
