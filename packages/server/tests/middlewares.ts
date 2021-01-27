import supertest from 'supertest';
import { getConnection } from 'typeorm';
import createTypeOrmConnection from '../dist/helpers/createTypeOrmConnection';
import app from '../dist/server';

const request = supertest(app);

beforeAll(async () => {
  await createTypeOrmConnection();
});

it('should return 404 when resource not found', async (done) => {
  const res = await request.get('/non-existing-resource').expect(404);
  expect(res.body).toHaveProperty('message');
  done();
});

afterAll(async () => {
  await getConnection().close();
});
