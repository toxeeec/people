import argon2 from 'argon2';
import faker from 'faker';
import supertest from 'supertest';
import { getConnection } from 'typeorm';
import { User } from '../../dist/entity/User';
import createTypeOrmConnection from '../../dist/helpers/createTypeOrmConnection';
import app from '../../dist/server';

const request = supertest(app);

interface IUser {
  name: string;
  surname: string;
  email: string;
  password?: string;
}

let dbUsers: User[];

beforeAll(async () => {
  await createTypeOrmConnection();
  const users: IUser[] = [];
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
  dbUsers = User.create(users);
  await User.save(dbUsers);
});

describe('requests route', () => {
  const path = '/api/requests';
  it('should send friend request to the user', async (done) => {
    const res = await request.post(`${path}/${dbUsers[1].id}`).expect(201);
    expect(res.body).toHaveProperty('message', 'Request sent successfully');
    done();
  });

  it('should return 400 error when request to this user was already sent', async (done) => {
    await request.post(`${path}/${dbUsers[1].id}`);
    const res = await request.post(`${path}/${dbUsers[1].id}`).expect(400);
    expect(res.body).toHaveProperty(
      'message',
      'Request to this user was already sent'
    );
    done();
  });
});

afterAll(async () => {
  await getConnection().close();
});
