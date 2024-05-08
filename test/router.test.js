const e = require("express");
const request = require("supertest");
describe("POST /say-name/:name", () => {
  let app;
  beforeAll(() => {
    app = e();
    app.listen(3000, () => {
      console.log("server run on port 300");
    });
  });
  afterAll(() => {
    app.close(() => {
      console.log("close server ");
    });
  });
  it("should return 201 when", async () => {
    const response = await request(app).post(`/say-name/mahdi`);
    console.log("response:", response);
    expect(response.body).toEqual({name:"mahdi"});
    expect(response.statusCode).toBe(200)

  });
});
