import argon2 from 'argon2';
import jwt from 'jsonwebtoken';
import supertest from 'supertest';
import { getConnection } from 'typeorm';
import { User } from '../../dist/entity/User';
import createTokens from '../../dist/helpers/createTokens';
import createTypeOrmConnection from '../../dist/helpers/createTypeOrmConnection';
import app from '../../dist/server';

const request = supertest(app);

const existingUser = {
  name: 'ExistingName',
  surname: 'ExistingSurname',
  email: 'existingemail@test.com',
  password: 'ExistingPassword123',
};

let user: User;

beforeAll(async () => {
  await createTypeOrmConnection();
  const { name, surname, email, password } = existingUser;
  const hashedPassword = await argon2.hash(password);
  user = await User.create({
    name,
    surname,
    email,
    password: hashedPassword,
  }).save();
});

describe('token route', () => {
  const path = '/api/token';

  it('should return 401 error when refresh token is not given', async (done) => {
    await request.post(path).expect(401);
    done();
  });

  it('should return 403 error when given invalidated refresh token', async (done) => {
    const refreshToken = jwt.sign(
      { id: user.id },
      process.env.REFRESH_TOKEN_SECRET!
    );
    await request.post(path).auth(refreshToken, { type: 'bearer' }).expect(403);
    done();
  });

  it('should return 403 error when given invalid refresh token', async (done) => {
    await request
      .post(path)
      .auth('refreshToken', { type: 'bearer' })
      .expect(403);
    done();
  });

  it('should return access token when given valid refresh token', async (done) => {
    const { refreshToken } = await createTokens(user);
    const res = await request
      .post(path)
      .auth(refreshToken, { type: 'bearer' })
      .expect(200);
    expect(res.body).toHaveProperty('accessToken');
    done();
  });
});

afterAll(async () => {
  await getConnection().close();
});
