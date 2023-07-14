import { PolywrapClient } from "@polywrap/client-js";
import { OsmoMath_Module } from "../types/wrap";
import path from "path";

jest.setTimeout(60000);

describe("Template Wrapper End to End Tests", () => {

  const client: PolywrapClient = new PolywrapClient();
  let wrapperUri: string;

  beforeAll(() => {
    const dirname: string = path.resolve(__dirname);
    const wrapperPath: string = path.join(dirname, "..", "..", "..");
    wrapperUri = `fs/${wrapperPath}/build`;
  })

  it("calls sampleMethod", async () => {
    const expected: string = "polywrap";

    const result = await client.invoke({
      uri: "file/build",
      method: "calcPriceImpactWithAmount",
      args: {
        spotPriceBefore: "15345",
        tokenAmount: "53234223",
        priceImpact: "12432"
      }
    });

    const result = await client.invoke<App.Template_SampleResult>({
      uri: wrapperUri,
      method: "sampleMethod",
      args: { arg: expected }
    });

    expect(result.ok).toBeTruthy();
    if (!result.ok) return;
    expect(result.value.result).toEqual(expected);
  });
});
