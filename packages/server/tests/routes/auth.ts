import argon2 from 'argon2';
import supertest from 'supertest';
import { getConnection } from 'typeorm';
import { User } from '../../dist/entity/User';
import createTypeOrmConnection from '../../dist/helpers/createTypeOrmConnection';
import app from '../../dist/server';

const request = supertest(app);

const testUser = {
  name: 'TestName',
  surname: 'TestSurname',
  email: 'test@test.com',
  password: 'TestPassword123',
};

const existingUser = {
  name: 'ExistingName',
  surname: 'ExistingSurname',
  email: 'existingemail@test.com',
  password: 'ExistingPassword123',
};
beforeAll(async () => {
  await createTypeOrmConnection();
  const { name, surname, email, password } = existingUser;
  const hashedPassword = await argon2.hash(password);
  await User.create({ name, surname, email, password: hashedPassword }).save();
});

describe('auth routes', () => {
  it('should return acess token when correct email and password', async (done) => {
    const res = await request
      .post('/api/login')
      .set('Accept', 'application/json')
      .send({
        email: existingUser.email,
        password: existingUser.password,
      })
      .expect(200);
    expect(res.body).toContainAllKeys([
      'accessToken',
      'name',
      'surname',
      'email',
    ]);
    done();
  });

  it('should return 401 error when wrong email', async (done) => {
    const res = await request
      .post('/api/login')
      .set('Accept', 'application/json')
      .send({
        email: 'wrongEmail@gmail.com',
        password: existingUser.password,
      })
      .expect(401);
    expect(res.body).toHaveProperty('message');
    done();
  });

  it('should return 401 error when wrong password', async (done) => {
    const res = await request
      .post('/api/login')
      .set('Accept', 'application/json')
      .send({
        email: existingUser.email,
        password: 'wrongPassword',
      })
      .expect(401);
    expect(res.body).toHaveProperty('message');
    done();
  });

  it('should return 400 error when required field is not passed', async (done) => {
    const res = await request
      .post('/api/login')
      .set('Accept', 'application/json')
      .send({
        email: existingUser.email,
      })
      .expect(400);
    expect(res.body).toHaveProperty('message');
    done();
  });

  it('should create user and send access token with user info when correct register credentials', async (done) => {
    const res = await request
      .post('/api/register')
      .set('Accept', 'application/json')
      .send(testUser)
      .expect(201);
    expect(res.body).toContainAllKeys([
      'accessToken',
      'name',
      'surname',
      'email',
    ]);
    done();
  });

  it('should return 400 error when wrong register credentials', async (done) => {
    const res = await request
      .post('/api/register')
      .set('accept', 'application/json')
      .send({ ...testUser, email: 'wrongemail' })
      .expect(400);
    expect(res.body).toHaveProperty('message');
    done();
  });

  it('should return 400 error when required field is not passed', async (done) => {
    const res = await request
      .post('/api/register')
      .set('accept', 'application/json')
      .send({ email: testUser.email, password: testUser.password })
      .expect(400);
    expect(res.body).toHaveProperty('message');
    done();
  });
});

afterAll(async () => {
  await getConnection().close();
});
