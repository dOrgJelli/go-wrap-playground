import { PolywrapClient } from "@polywrap/client-js";

async function main() {
  const client = new PolywrapClient();

  const result = await client.invoke({
    uri: "file/build",
    method: "calcPriceImpactWithAmount",
    args: {
      spotPriceBefore: "15345",
      tokenAmount: "53234223",
      priceImpact: "12432"
    }
  });

  console.log(result);
}

main();
