import e, { Application } from "express";
import request from "supertest";
import { afterAll, beforeAll, describe, expect, it } from "@jest/globals";
describe("POST /say-name/:name", () => {
  let app: Application, server: any;
  beforeAll(() => {
    app = e();
    server = app.listen(3000);
  });
  afterAll(() => {
    server.close();
  });
  it("should return 201 Created with the name in the response body", async () => {
    const response = await request(app).post("/say-name/mahdi");
    expect(response.statusCode).toBe(201);
    expect(response.body).toEqual({ name: "mahdi" });
  });
});
