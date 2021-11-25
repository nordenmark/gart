import React from "react";

import { ImageParameter } from "./interfaces";

interface ControlsProps {
  config: { [key: string]: string | number };
  parameters: ImageParameter[];
  onControlChanged: (name: string, value: string | number) => void;
}

interface ControlProps {
  value: string | number;
  parameter: ImageParameter;
  onControlChanged: (name: string, value: string | number) => void;
}

function NumberControl({
  value,
  parameter: { name, min, max },
  onControlChanged,
}: ControlProps) {
  min = min ?? 0;
  max = max ?? 99999;

  function onChange(e: React.ChangeEvent<HTMLInputElement>) {
    if (!e.target.validity.valid) {
      return;
    }

    const value = parseInt(e.target.value, 10);

    onControlChanged(name, value);
  }

  return (
    <div>
      <label>{name}</label>
      <input
        id={name}
        type="number"
        min={min}
        max={max}
        value={value}
        onChange={onChange}
      />
    </div>
  );
}

function RangeControl({
  value,
  parameter: { name, min, max, step },
  onControlChanged,
}: ControlProps) {
  return (
    <div>
      <label>{name}</label>
      <input
        type="range"
        min={min}
        max={max}
        step={step}
        value={value}
        onChange={(e) => onControlChanged(name, parseFloat(e.target.value))}
      />
      <span>{value}</span>
    </div>
  );
}

function Controls({ parameters, config, onControlChanged }: ControlsProps) {
  const controls = parameters.map((parameter) => {
    switch (parameter.type) {
      case "number":
        return (
          <NumberControl
            key={parameter.name}
            value={config[parameter.name]}
            parameter={parameter}
            onControlChanged={onControlChanged}
          ></NumberControl>
        );
      case "range":
        return (
          <RangeControl
            key={parameter.name}
            value={config[parameter.name]}
            parameter={parameter}
            onControlChanged={onControlChanged}
          ></RangeControl>
        );
      default:
        return "";
    }
  });

  return <div>{controls}</div>;
}

export default Controls;
