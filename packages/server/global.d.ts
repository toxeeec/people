import 'express';
import 'jest-extended';
import { EntityTarget } from 'typeorm';

declare module 'express' {
  export interface Request {
    user?: {
      id: string;
      name: string;
      surname: string;
    };
    entity?: EntityTarget;
  }
}
