import argon2 from 'argon2';
import faker from 'faker';
import supertest from 'supertest';
import { getConnection } from 'typeorm';
import { User } from '../../../dist/entity/User';
import createTokens from '../../../dist/helpers/createTokens';
import createTypeOrmConnection from '../../../dist/helpers/createTypeOrmConnection';
import app from '../../../dist/server';

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
  for (let i = 0; i < 6; i += 1) {
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

describe('requests route get', () => {
  const path = '/api/requests';
  it('should return sent friend requests', async (done) => {
    for await (const i of [1, 2]) {
      await request
        .post(`${path}/${users[i].id}`)
        .auth(users[0].accessToken!, { type: 'bearer' });
    }
    const res = await request
      .get(`${path}/sent`)
      .auth(users[0].accessToken!, { type: 'bearer' })
      .expect(200);
    expect(res.body).toHaveProperty('count', 2);
    done();
  });

  it('should return received friend requests', async (done) => {
    for await (const i of [3, 4]) {
      await request
        .post(`${path}/${users[5].id}`)
        .auth(users[i].accessToken!, { type: 'bearer' });
    }
    const res = await request
      .get(`${path}/received`)
      .auth(users[5].accessToken!, { type: 'bearer' })
      .expect(200);
    expect(res.body).toHaveProperty('count', 2);
    done();
  });
});

afterAll(async () => {
  await getConnection().close();
});
