import argon2 from 'argon2';
import faker from 'faker';
import supertest from 'supertest';
import { getConnection } from 'typeorm';
import { User } from '../../../src/entity/User';
import createTokens from '../../../src/helpers/createTokens';
import createTypeOrmConnection from '../../../src/helpers/createTypeOrmConnection';
import app from '../../../src/server';

const request = supertest(app);

interface IUser {
  id?: string;
  name: string;
  surname: string;
  email: string;
  password?: string;
  accessToken?: string;
}

const users: IUser[] = [];

beforeAll(async () => {
  await createTypeOrmConnection();
  for (let i = 0; i < 2; i += 1) {
    users.push({
      name: faker.name.firstName(),
      surname: faker.name.lastName(),
      email: faker.internet.email(),
    });
  }
  for await (const user of users) {
    user.password = await argon2.hash(faker.internet.password());
  }
  let dbUsers = User.create(users);
  dbUsers = await User.save(dbUsers);
  for await (const [i, user] of users.entries()) {
    user.id = dbUsers[i].id;
    const { accessToken } = await createTokens(user as User);
    user.accessToken = accessToken;
  }
});

describe('requests route delete', () => {
  const path = '/api/requests';
  it('should delete sent friend request', async (done) => {
    await request
      .post(`${path}/${users[1].id}`)
      .auth(users[0].accessToken!, { type: 'bearer' });
    const res = await request
      .delete(`${path}/${users[1].id}`)
      .auth(users[0].accessToken!, { type: 'bearer' })
      .expect(200);
    expect(res.body).toHaveProperty('message', 'Request deleted successfully');
    done();
  });
});

afterAll(async () => {
  await getConnection().close();
});
