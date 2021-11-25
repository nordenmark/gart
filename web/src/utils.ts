import { ImageParameter } from "./interfaces";

export function buildImageUrl(
  imageName: string,
  config: { [key: string]: unknown }
) {
  const base = "http://localhost:8080";
  const url = new URL(`/g/${imageName}`, base);

  Object.entries(config).forEach(([key, value]) => {
    url.searchParams.append(key, value as string);
  });

  // Cache bust
  url.searchParams.append("t", Date.now().toString());

  return url.toString();
}

export function getDefaultConfigValues(parameters: ImageParameter[]): {
  [key: string]: unknown;
} {
  return parameters.reduce((acc, curr) => {
    return {
      ...acc,
      [curr.name]: curr.defaultValue,
    };
  }, {});
}
