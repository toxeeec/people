import supertest from 'supertest';
import { createConnection, getConnection } from 'typeorm';
import app from '../../dist/server';

const { NODE_ENV } = process.env;

const request = supertest(app);

const testUser = {
  name: 'TestName',
  surname: 'TestSurname',
  email: 'test@test.com',
  password: 'TestPassword123',
};
beforeAll(async () => {
  await createConnection(NODE_ENV!);
});

describe('auth routes', () => {
  it('should return acess token when correct email and password', async (done) => {
    const res = await request
      .post('/login')
      .set('Accept', 'application/json')
      .send({
        email: testUser.email,
        password: testUser.password,
      })
      .expect(200);
    expect(res.body).toHaveProperty('accessToken');
    done();
  });

  it('should return 401 error when wrong email', async (done) => {
    const res = await request
      .post('/login')
      .set('Accept', 'application/json')
      .send({
        email: 'wrongEmail',
        password: testUser.password,
      })
      .expect(401);
    expect(res.body).toHaveProperty('message');
    done();
  });

  it('should return 401 error when wrong password', async (done) => {
    const res = await request
      .post('/login')
      .set('Accept', 'application/json')
      .send({
        email: testUser.email,
        password: 'wrongPassword',
      })
      .expect(401);
    expect(res.body).toHaveProperty('message');
    done();
  });

  it('should return 400 error when required field is not passed', async (done) => {
    const res = await request
      .post('/login')
      .set('Accept', 'application/json')
      .send({
        email: testUser.email,
      })
      .expect(401);
    expect(res.body).toHaveProperty('message');
    done();
  });

  it('should create user and send access token when correct register credentials', async (done) => {
    const res = await request
      .post('/register')
      .set('Accept', 'application/json')
      .send(testUser)
      .expect(201);
    expect(res.body).toHaveProperty('accessToken');
    done();
  });

  it('should return 400 error when wrong register credentials', async (done) => {
    const res = await request
      .post('/register')
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
  await getConnection(NODE_ENV).close();
});
