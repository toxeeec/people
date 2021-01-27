import 'express';
import 'jest-extended';

declare module 'express' {
  export interface Request {
    user?: {
      id: number;
    };
  }
}
