import axios from "axios";

const invInstances = [
  "https://invidio.xamh.de",
  "https://invidious.snopyta.org",
];
const channals = [
  { name: "Jonathan Pageau", channelID: "UCtCTSf3UwRU14nYWr_xm-dQ" },
];
const videoIDs = []; // if you want to include specific video but not the whole channel

const channalENDPOINT = "/api/v1/channels/videos/"; // + :uid
const videosENDPOINT = "/api/v1/videos/"; // + :id

// this should run only once a month, otherwise get videos from cache, and maybe move this logic to GO
export const getVideoUrls = async () => {
  let videoIDs = await getAllVideosFromChannel(channals[0].channelID);
  let videoURLs = await getUrlByID(videoIDs);
  // save videourls
  return videoURLs;
};

const getAllVideosFromChannel = async (channelID) => {
  // with while iterate over every instance
  let res = await axios.get(invInstances[0] + channalENDPOINT + channelID);
  return res.data.map((info) => info.videoId);
};

const getUrlByID = async (videoIDs) => {
  let videoUrls = [];
  for (let i = 0; i < videoIDs.length; i++) {
    let res = await axios.get(invInstances[0] + videosENDPOINT + videoIDs[i]);
    let videoTypes = res.data.formatStreams;
    let url = videoTypes.find((el) => el.itag === "18").url;
    if (url) videoUrls.push(url);
  }
  return videoUrls;
};
