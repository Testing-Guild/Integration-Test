const e = require("express");
const request = require("supertest");
describe("POST /say-name/:name", () => {
  let app;
  let server; 

  beforeAll(() => {
    app = e();
    app.post('/say-name/:name', (req, res) => {
      const { name } = req.params;
      res.status(201).json({ name }); 
    });

    server = app.listen(3000);
  });

  afterAll(() => {
    server.close();
  });
  it('should return 201 Created with the name in the response body', async () => {
    const response = await request(app).post('/say-name/mahdi');
    expect(response.statusCode).toBe(201); 
    expect(response.body).toEqual({ name: 'mahdi' });
  });
});
