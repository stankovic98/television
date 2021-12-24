import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const VideoBox = ({ video }) => {
  useEffect(() => {
    // console.log(video);
  }, []);
  const navigate = useNavigate();
  const playVideo = () => {
    navigate("/watch", { state: { id: video.ID, title: video.Name } });
  };
  return (
    <div className="videoBox" onClick={playVideo}>
      <h3>{video.Name}</h3>
      <img src={video.ThumbnailUrl} />
    </div>
  );
};

export default VideoBox;
