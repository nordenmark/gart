import React, { useState, useEffect } from "react";

import {
  Container,
  GenerateBtn,
  Header,
  ImageContainer,
  ImageLink,
  ImageList,
  Main,
  OutputImage,
  Sidebar,
  Spinner,
} from "./App.styles";
import { API_HOST } from "./configuration";
import Controls from "./Controls";
import { ImageObject } from "./interfaces";
import spinner from "./spinner.svg";
import { buildImageUrl, getDefaultConfigValues } from "./utils";

function App() {
  const [images, setImages] = useState<ImageObject[]>([]);
  const [selectedImage, setSelectedImage] = useState<ImageObject | null>(null);
  const [imageUrl, setImageUrl] = useState<string | null>(null);
  const [config, setConfig] = useState({});
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    fetch(`${API_HOST}/api/images`)
      .then((res) => res.json())
      .then((images: ImageObject[]) => {
        setImages(images);
      });
  }, []);

  function handleGenerate() {
    if (!selectedImage) {
      return;
    }

    const url = buildImageUrl(selectedImage.name, config);

    setLoading(true);
    setImageUrl(url);
  }

  function handleImageLoaded() {
    setLoading(false);
  }

  function handleImageSelected(image: ImageObject) {
    setSelectedImage(image);

    const defaultValues = getDefaultConfigValues(image.parameters);

    setConfig(defaultValues);
  }

  function handleControlChanged(name: string, value: string | number) {
    setConfig({
      ...config,
      [name]: value,
    });
  }

  const listItems = images.map((image) => (
    <li key={image.name}>
      <ImageLink
        className={image.name === selectedImage?.name ? "active" : ""}
        onClick={() => handleImageSelected(image)}
      >
        {image.name}
      </ImageLink>
    </li>
  ));

  const outputImage = imageUrl ? (
    <OutputImage onLoad={handleImageLoaded} src={imageUrl}></OutputImage>
  ) : (
    ""
  );

  const controls = selectedImage ? (
    <Controls
      config={config}
      parameters={selectedImage.parameters}
      onControlChanged={handleControlChanged}
    ></Controls>
  ) : (
    ""
  );

  const loadingIndicator = loading ? <Spinner src={spinner}></Spinner> : "";

  return (
    <Container>
      <Sidebar>
        <h3>gART</h3>
        <ImageList>{listItems}</ImageList>
        <GenerateBtn
          disabled={!selectedImage || loading}
          onClick={() => handleGenerate()}
        >
          Generate
        </GenerateBtn>
        {loadingIndicator}
      </Sidebar>
      <Main>
        <Header>{controls}</Header>
        <ImageContainer>{outputImage}</ImageContainer>
      </Main>
    </Container>
  );
}

export default App;
