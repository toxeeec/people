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
  for (let i = 0; i < 3; i += 1) {
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

describe('requests route patch', () => {
  const path = '/api/requests';
  it('should accept the friend request', async (done) => {
    await request
      .post(`${path}/${users[0].id}`)
      .auth(users[2].accessToken!, { type: 'bearer' });
    const res = await request
      .patch(`${path}/${users[2].id}`)
      .auth(users[0].accessToken!, { type: 'bearer' })
      .set('Accept', 'application/json')
      .send({ action: 'accept' })
      .expect(200);
    expect(res.body).toHaveProperty('message', 'Request accepted successfully');
    done();
  });

  it('should decline the friend request', async (done) => {
    await request
      .post(`${path}/${users[1].id}`)
      .auth(users[2].accessToken!, { type: 'bearer' });
    const res = await request
      .patch(`${path}/${users[2].id}`)
      .auth(users[1].accessToken!, { type: 'bearer' })
      .set('Accept', 'application/json')
      .send({ action: 'decline' })
      .expect(200);
    expect(res.body).toHaveProperty('message', 'Request declined successfully');
    done();
  });

  it('should return 400 error when action is not accept or decline', async (done) => {
    const res = await request
      .patch(`${path}/${users[2].id}`)
      .auth(users[0].accessToken!, { type: 'bearer' })
      .set('Accept', 'application/json')
      .send({ action: 'wrong' })
      .expect(400);
    expect(res.body).toHaveProperty(
      'message',
      'action must be one of the following values: accept, decline'
    );
    done();
  });
});

afterAll(async () => {
  await getConnection().close();
});
