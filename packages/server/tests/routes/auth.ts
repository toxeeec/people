import argon2 from 'argon2';
import faker from 'faker';
import supertest from 'supertest';
import { getConnection } from 'typeorm';
import { User } from '../../src/entity/User';
import createTypeOrmConnection from '../../src/helpers/createTypeOrmConnection';
import app from '../../src/server';

const request = supertest(app);

const sampleUser = {
  name: faker.name.firstName(),
  surname: faker.name.lastName(),
  email: faker.internet.email(),
  password: faker.internet.password(),
};

const existingUser = {
  name: faker.name.firstName(),
  surname: faker.name.lastName(),
  email: faker.internet.email(),
  password: faker.internet.password(),
};

beforeAll(async () => {
  await createTypeOrmConnection();
  const { name, surname, email, password } = existingUser;
  const hashedPassword = await argon2.hash(password);
  await User.create({
    name,
    surname,
    email,
    password: hashedPassword,
  }).save();
});

describe('login route', () => {
  const path = '/api/login';

  it('should return user info and tokens when given correct email and password', async (done) => {
    const res = await request
      .post(path)
      .set('Accept', 'application/json')
      .send({
        email: existingUser.email,
        password: existingUser.password,
      })
      .expect(200);
    expect(res.body).toContainAllKeys([
      'name',
      'surname',
      'email',
      'refreshToken',
      'accessToken',
    ]);
    done();
  });

  it('should return 401 error when given wrong email', async (done) => {
    await request
      .post(path)
      .set('Accept', 'application/json')
      .send({
        email: 'wrongEmail@gmail.com',
        password: existingUser.password,
      })
      .expect(401);
    done();
  });

  it('should return 401 error when given wrong password', async (done) => {
    await request
      .post(path)
      .set('Accept', 'application/json')
      .send({
        email: existingUser.email,
        password: 'wrongPassword',
      })
      .expect(401);
    done();
  });
});

describe('register route', () => {
  const path = '/api/register';

  it('should create user, return user info and tokens when given correct register credentials', async (done) => {
    const res = await request
      .post(path)
      .set('Accept', 'application/json')
      .send(sampleUser)
      .expect(201);
    expect(res.body).toContainAllKeys([
      'name',
      'surname',
      'email',
      'refreshToken',
      'accessToken',
    ]);
    done();
  });

  it('should return 400 error when given email is not unique', async (done) => {
    await request
      .post(path)
      .set('Accept', 'application/json')
      .send(sampleUser)
      .expect(400);
    done();
  });
});

afterAll(async () => {
  await getConnection().close();
});
