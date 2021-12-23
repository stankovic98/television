import axios from "axios";
import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { getVideoUrls } from "../services/InvidiusAPI";

const invInstances = [
  "https://invidio.xamh.de",
  "https://invidious.snopyta.org",
];

const videoEndpoint = "/api/v1/videos/"; // + videoID

// if this still works after 2 days, we can procide with the work (if there is no expire or access denaied)
const WatchVideo = () => {
  const { state } = useLocation();
  const [url, setUrl] = useState();

  const getVideoUrl = async () => {
    let res = await axios.get(invInstances[0] + videoEndpoint + state.id);
    let videoTypes = res.data.formatStreams;
    let url = videoTypes.find((el) => el.itag === "18").url;
    setUrl(url);
    console.log(url);
  };

  useEffect(() => {
    getVideoUrl();
  }, []);

  return (
    <div>
      <h3>Title {state.id}</h3>
      {url && (
        <video
          width="400"
          height="200"
          controls
          disablePictureInPicture
          autoPlay
        >
          <source src={url} type="video/mp4" />
          Your browser does not support the video tag.
        </video>
      )}
    </div>
  );
};

export default WatchVideo;
