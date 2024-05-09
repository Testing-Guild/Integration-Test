import express, {
  Express,
  RouterOptions,
  IRouter,
  Request,
  Response,
} from "express";
class SayMyNam {
  name: string;
  constructor(name: string) {
    this.name = name;
  }
}
const router: IRouter = express();
router.post("/say-name/:name", (req: Request, res: Response) => {
  const { name } = req.params;
  res.send({ name: new SayMyNam(name) });
});

export default router;
