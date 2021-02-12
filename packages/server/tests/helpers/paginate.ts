import argon2 from 'argon2';
import { NextFunction } from 'express';
import faker from 'faker';
import { User } from '../../dist/entity/User';
import createTypeOrmConnection from '../../dist/helpers/createTypeOrmConnection';
import paginate from '../../dist/helpers/paginate';

interface IUser {
  name: string;
  surname: string;
  email: string;
  password?: string;
}

describe('validate helper function', () => {
  const nextFunction: NextFunction = jest.fn();

  beforeAll(async () => {
    await createTypeOrmConnection();
    const users: IUser[] = [];
    for (let i = 0; i < 40; i += 1) {
      users.push({
        name: faker.name.firstName(),
        surname: faker.name.lastName(),
        email: faker.internet.email(),
      });
    }
    for await (const user of users) {
      user.password = await argon2.hash(faker.internet.password());
    }
    const dbUsers = User.create(users);
    await User.save(dbUsers);
  });

  it('should return results = limit and next parameter when given page is first and limit < length of query results', async (done) => {
    const res = await paginate(nextFunction, User, { page: 1, limit: 20 });
    expect(res).toHaveProperty('count', 40);
    expect(res).toHaveProperty('next');
    expect(res.data).toHaveLength(20);
    done();
  });

  it('should return results = limit and next, previous parameters when given page is in the middle and limit < length of query results', async (done) => {
    const res = await paginate(nextFunction, User, { page: 2, limit: 10 });
    expect(res).toHaveProperty('count', 40);
    expect(res).toContainKeys(['previous', 'next']);
    expect(res.data).toHaveLength(10);
    done();
  });

  it('should return results = limit and previous parameter when given page is last and limit < length of query results', async (done) => {
    const res = await paginate(nextFunction, User, { page: 2, limit: 20 });
    expect(res).toHaveProperty('count', 40);
    expect(res).toHaveProperty('previous');
    expect(res.data).toHaveLength(20);
    done();
  });

  it('should return results = remaining query results when page is not first and limit < length of query results', async (done) => {
    const res = await paginate(nextFunction, User, { page: 3, limit: 15 });
    expect(res).toHaveProperty('count', 40);
    expect(res.data).toHaveLength(10);
    done();
  });

  it('should return results = length of query results when page is first and limit > length of query results', async (done) => {
    const res = await paginate(nextFunction, User, { page: 1, limit: 50 });
    expect(res).toHaveProperty('count', 40);
    expect(res.data).toHaveLength(40);
    done();
  });

  it('should return results = length of query results when limit is not given', async (done) => {
    const res = await paginate(nextFunction, User);
    expect(res).toHaveProperty('count', 40);
    expect(res.data).toHaveLength(40);
    done();
  });

  it('should return empty results and previous parameter when given page is out of range', async (done) => {
    const res = await paginate(nextFunction, User, { page: 5, limit: 20 });
    expect(res).toHaveProperty('count', 40);
    expect(res).toHaveProperty('previous');
    expect(res.data).toHaveLength(0);
    done();
  });
});
