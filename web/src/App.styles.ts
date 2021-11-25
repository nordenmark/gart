import styled from "styled-components";

export const Container = styled.div`
  display: flex;
  height: 100%;
`;

export const Sidebar = styled.aside`
  background-color: rgb(71, 150, 220);
  padding: 10px;
  color: white;
  height: 100%;
  width: 200px;
  display: flex;
  flex-direction: column;
`;

export const GenerateBtn = styled.button`
  padding: 10px 20px;
  width: 100%;
  margin: 0;
  font-size: 20px;
  &:hover {
    cursor: pointer;
  }
`;

export const ImageList = styled.ul`
  margin-top: 20px;
  margin-bottom: 20px;
`;

export const ImageLink = styled.a`
  color: white;
  font-size: 20px;
  margin-bottom: 8px;
  display: block;
  &.active {
    text-decoration: underline;
  }
`;

export const Main = styled.main`
  width: calc(100% - 200px);
  display: flex;
  flex-direction: column;
`;

export const Header = styled.header`
  padding: 10px 30px;
  min-height: 82px;
`;

export const ImageContainer = styled.div`
  background: beige;
  width: 100%;
  height: 100%;
  padding: 30px;
`;

export const OutputImage = styled.img`
  width: 50%;
  height: auto;
`;

export const Spinner = styled.img`
  margin: 20px auto 0 auto;
`;
