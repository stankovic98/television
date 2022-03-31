import axios from "axios";
import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
// import { getVideoUrls } from "../services/InvidiusAPI";
import { useNavigate } from "react-router-dom";

const invInstances = [
  "https://invidio.xamh.de",
  "https://invidious.snopyta.org",
];

const videoEndpoint = "/api/v1/videos/"; // + videoID

const WatchVideo = () => {
  const { state } = useLocation();
  const navigate = useNavigate();
  const [url, setUrl] = useState();
  const [isLoading, setIsLoading] = useState(true);

  const getVideoUrl = async () => {
    let res = await axios.get(invInstances[0] + videoEndpoint + state.id);
    let videoTypes = res.data.formatStreams;
    let url = videoTypes.find((el) => el.itag === "18").url;
    setUrl(url);
    setIsLoading(false);
    console.log(url);
  };

  useEffect(() => {
    getVideoUrl();
  }, []);

  return (
    <div className="watch-video-comp">
      <button className="back" onClick={() => navigate("/feed")}>
        Back to Feed
      </button>
      <h3 className="title">{state.title}</h3>
      {isLoading && <div className="lds-dual-ring"></div>}
      {url && (
        <video
          controls
          disablePictureInPicture
          autoPlay
          className="video-player"
        >
          <source src={url} type="video/mp4" />
          Your browser does not support the video tag.
        </video>
      )}
    </div>
  );
};

export default WatchVideo;
