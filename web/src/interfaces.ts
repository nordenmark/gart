export interface ImageParameter {
  name: string;
  type: "number" | "range";
  defaultValue: number;
  min: number;
  max: number;
  step: number;
}

export interface ImageObject {
  name: string;
  parameters: ImageParameter[];
}
