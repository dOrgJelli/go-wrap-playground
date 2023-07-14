import { PolywrapClient } from "@polywrap/client-js";
import { OsmoMath_Module as OsmoMath } from "../types/wrap";
import path from "path";

jest.setTimeout(60000);

describe("Osmosis Math End to End Tests", () => {

  const client: PolywrapClient = new PolywrapClient();
  const wrapUri = `file/${path.join(__dirname, "../../../build")}`;

  it("calcPriceImpactWithAmount", async () => {

    const result = await OsmoMath.calcPriceImpactWithAmount({
      spotPriceBefore: "15345",
      tokenAmount: "53234223",
      priceImpact: "12432"
    }, client, wrapUri);

    expect(result.ok).toBeTruthy();
    if (!result.ok) return;
    expect(result.value).toEqual("0.279028197197585117");
  });
});
