import argon2 from 'argon2';
import dotenv from 'dotenv';
import { NextFunction, Request, Response } from 'express';
import faker from 'faker';
import { getConnection } from 'typeorm';
import { User } from '../../dist/entity/User';
import ApiError from '../../dist/helpers/ApiError';
import createTypeOrmConnection from '../../dist/helpers/createTypeOrmConnection';
import getEntityById from '../../dist/middlewares/getEntityById';

dotenv.config({ path: '.env.test' });

describe('getEntityById middleware', () => {
  let mockRequest: Partial<Request>;
  let mockResponse: Partial<Response>;
  const nextFunction: NextFunction = jest.fn();
  let userId: string;

  beforeAll(async () => {
    await createTypeOrmConnection();
    const sampleUser = {
      name: faker.name.firstName(),
      surname: faker.name.lastName(),
      email: faker.internet.email(),
      password: await argon2.hash(faker.internet.password()),
    };
    const { id } = await User.create(sampleUser).save();
    userId = id;
  });

  beforeEach(() => {
    mockRequest = { params: {}, entity: {} };
    mockResponse = {};
  });

  it('should call next function and return entity in request if user exists', async (done) => {
    mockRequest.params!.id = userId;
    await getEntityById('User')(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(mockRequest.entity).toContainAllKeys([
      'id',
      'name',
      'surname',
      'email',
      'password',
    ]);
    expect(nextFunction).toHaveBeenCalledWith();
    done();
  });

  it('should call next function with 404 ApiError error when user does not exist', async (done) => {
    mockRequest.params!.id = faker.random.uuid();
    await getEntityById('User')(
      mockRequest as Request,
      mockResponse as Response,
      nextFunction
    );
    expect(nextFunction).toHaveBeenLastCalledWith(ApiError.notFound());
    done();
  });
  afterAll(async () => {
    await getConnection().close();
  });
});
