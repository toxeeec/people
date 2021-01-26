import supertest from 'supertest';
import app from '../../dist/index';

const request = supertest(app);

describe('auth routes', () => {
  it('should return acess token when correct login credentials', async () => {
    const res = await request.post('/login');
    expect(res.body).toHaveProperty('accessToken');
  });
});
