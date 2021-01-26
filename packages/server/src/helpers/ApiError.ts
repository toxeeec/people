export class ApiError {
  status: number;

  message: string;

  constructor(status: number, message: string) {
    this.status = status;
    this.message = message;
  }

  static 401() {
    return new ApiError(401, 'Unauthorized');
  }

  static 403() {
    return new ApiError(403, 'Forbidden');
  }

  static 404() {
    return new ApiError(404, 'Not Found');
  }

  static 500() {
    return new ApiError(500, 'Internal Server Error');
  }
}
