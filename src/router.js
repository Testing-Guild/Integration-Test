const express = require("express");
class SayMyNam {
  constructor(name) {
    this.name = name;
  }
}
const router = express.Router();

router.post("/say-name/:name", (req, res) => {
  const { name } = req.params;
  res.send({ name: new SayMyNam(name) });
});

module.exports = router;
