import supertest from 'supertest';
import { createConnection, getConnection } from 'typeorm';
import app from '../dist/server';

const request = supertest(app);

beforeAll(async () => {
  await createConnection();
});

it('should return 404 when resource not found', async (done) => {
  const res = await request.get('/non-existing-resource').expect(404);
  expect(res.body).toHaveProperty('message');
  done();
});

afterAll(async () => {
  await getConnection().close();
});
