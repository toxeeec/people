import argon2 from 'argon2';
import faker from 'faker';
import supertest from 'supertest';
import { getConnection } from 'typeorm';
import { User } from '../../dist/entity/User';
import createTypeOrmConnection from '../../dist/helpers/createTypeOrmConnection';
import app from '../../dist/server';

const request = supertest(app);

const existingUser = {
  name: faker.name.firstName(),
  surname: faker.name.lastName(),
  email: faker.internet.email(),
  password: faker.internet.password(),
};

let id: string;

beforeAll(async () => {
  await createTypeOrmConnection();
  const { name, surname, email, password } = existingUser;
  const hashedPassword = await argon2.hash(password);
  const user = await User.create({
    name,
    surname,
    email,
    password: hashedPassword,
  }).save();
  id = user.id;
});

describe('auth routes', () => {
  const path = '/api/users';

  it('should return user info when given correct id and user exists', async (done) => {
    const res = await request.get(`${path}/${id}`).expect(200);
    expect(res.body).toContainAllKeys(['id', 'name', 'surname']);
    done();
  });
});

afterAll(async () => {
  await getConnection().close();
});
