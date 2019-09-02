package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/robfig/cron"
)

type Weibo struct {
	Ok   int `json:"ok"`
	Data struct {
		CardlistInfo struct {
			Containerid string `json:"containerid"`
			VP          int    `json:"v_p"`
			ShowStyle   int    `json:"show_style"`
			Total       int    `json:"total"`
			Page        int    `json:"page"`
		} `json:"cardlistInfo"`
		Cards []struct {
			CardType int    `json:"card_type"`
			Itemid   string `json:"itemid"`
			Scheme   string `json:"scheme"`
			Mblog    struct {
				CreatedAt                string `json:"created_at"`
				ID                       string `json:"id"`
				Idstr                    string `json:"idstr"`
				Mid                      string `json:"mid"`
				CanEdit                  bool   `json:"can_edit"`
				ShowAdditionalIndication int    `json:"show_additional_indication"`
				Text                     string `json:"text"`
				TextLength               int    `json:"textLength"`
				Source                   string `json:"source"`
				Favorited                bool   `json:"favorited"`
				PicTypes                 string `json:"pic_types"`
				ThumbnailPic             string `json:"thumbnail_pic"`
				BmiddlePic               string `json:"bmiddle_pic"`
				OriginalPic              string `json:"original_pic"`
				IsPaid                   bool   `json:"is_paid"`
				MblogVipType             int    `json:"mblog_vip_type"`
				User                     struct {
					ID              int64  `json:"id"`
					ScreenName      string `json:"screen_name"`
					ProfileImageURL string `json:"profile_image_url"`
					ProfileURL      string `json:"profile_url"`
					StatusesCount   int    `json:"statuses_count"`
					Verified        bool   `json:"verified"`
					VerifiedType    int    `json:"verified_type"`
					VerifiedTypeExt int    `json:"verified_type_ext"`
					VerifiedReason  string `json:"verified_reason"`
					CloseBlueV      bool   `json:"close_blue_v"`
					Description     string `json:"description"`
					Gender          string `json:"gender"`
					Mbtype          int    `json:"mbtype"`
					Urank           int    `json:"urank"`
					Mbrank          int    `json:"mbrank"`
					FollowMe        bool   `json:"follow_me"`
					Following       bool   `json:"following"`
					FollowersCount  int    `json:"followers_count"`
					FollowCount     int    `json:"follow_count"`
					CoverImagePhone string `json:"cover_image_phone"`
					AvatarHd        string `json:"avatar_hd"`
					Like            bool   `json:"like"`
					LikeMe          bool   `json:"like_me"`
					Badge           struct {
						Dzwbqlx2016         int `json:"dzwbqlx_2016"`
						UserNameCertificate int `json:"user_name_certificate"`
					} `json:"badge"`
				} `json:"user"`
				RepostsCount         int  `json:"reposts_count"`
				CommentsCount        int  `json:"comments_count"`
				AttitudesCount       int  `json:"attitudes_count"`
				PendingApprovalCount int  `json:"pending_approval_count"`
				IsLongText           bool `json:"isLongText"`
				RewardExhibitionType int  `json:"reward_exhibition_type"`
				HideFlag             int  `json:"hide_flag"`
				Visible              struct {
					Type   int `json:"type"`
					ListID int `json:"list_id"`
				} `json:"visible"`
				Mblogtype             int `json:"mblogtype"`
				MoreInfoType          int `json:"more_info_type"`
				ExternSafe            int `json:"extern_safe"`
				NumberDisplayStrategy struct {
					ApplyScenarioFlag    int    `json:"apply_scenario_flag"`
					DisplayTextMinNumber int    `json:"display_text_min_number"`
					DisplayText          string `json:"display_text"`
				} `json:"number_display_strategy"`
				ContentAuth       int `json:"content_auth"`
				PicNum            int `json:"pic_num"`
				MblogMenuNewStyle int `json:"mblog_menu_new_style"`
				EditConfig        struct {
					Edited bool `json:"edited"`
				} `json:"edit_config"`
				IsTop           int    `json:"isTop"`
				WeiboPosition   int    `json:"weibo_position"`
				ShowAttitudeBar int    `json:"show_attitude_bar"`
				ObjExt          string `json:"obj_ext"`
				PageInfo        struct {
					PagePic struct {
						URL string `json:"url"`
					} `json:"page_pic"`
					PageURL   string `json:"page_url"`
					PageTitle string `json:"page_title"`
					Content1  string `json:"content1"`
					Content2  string `json:"content2"`
					Type      string `json:"type"`
					MediaInfo struct {
						VideoOrientation   string `json:"video_orientation"`
						Name               string `json:"name"`
						StreamURL          string `json:"stream_url"`
						StreamURLHd        string `json:"stream_url_hd"`
						H5URL              string `json:"h5_url"`
						Mp4SdURL           string `json:"mp4_sd_url"`
						Mp4HdURL           string `json:"mp4_hd_url"`
						H265Mp4Hd          string `json:"h265_mp4_hd"`
						H265Mp4Ld          string `json:"h265_mp4_ld"`
						Inch4Mp4Hd         string `json:"inch_4_mp4_hd"`
						Inch5Mp4Hd         string `json:"inch_5_mp4_hd"`
						Inch55Mp4Hd        string `json:"inch_5_5_mp4_hd"`
						Mp4720PMp4         string `json:"mp4_720p_mp4"`
						HevcMp4720P        string `json:"hevc_mp4_720p"`
						PrefetchType       int    `json:"prefetch_type"`
						PrefetchSize       int    `json:"prefetch_size"`
						ActStatus          int    `json:"act_status"`
						Protocol           string `json:"protocol"`
						MediaID            string `json:"media_id"`
						OriginTotalBitrate int    `json:"origin_total_bitrate"`
						Duration           int    `json:"duration"`
						NextTitle          string `json:"next_title"`
						VideoDetails       []struct {
							Size         int    `json:"size"`
							Bitrate      int    `json:"bitrate"`
							Label        string `json:"label"`
							PrefetchSize int    `json:"prefetch_size"`
						} `json:"video_details"`
						HevcMp4Hd             string `json:"hevc_mp4_hd"`
						PlayCompletionActions []struct {
							Type         string `json:"type"`
							Icon         string `json:"icon"`
							Text         string `json:"text"`
							Link         string `json:"link"`
							BtnCode      int    `json:"btn_code"`
							ShowPosition int    `json:"show_position"`
							Actionlog    struct {
								Oid     string `json:"oid"`
								ActCode int    `json:"act_code"`
								ActType int    `json:"act_type"`
								Source  string `json:"source"`
							} `json:"actionlog"`
						} `json:"play_completion_actions"`
						VideoPublishTime int `json:"video_publish_time"`
						PlayLoopType     int `json:"play_loop_type"`
						Titles           []struct {
							Default bool   `json:"default"`
							Title   string `json:"title"`
						} `json:"titles"`
						AuthorMid      string `json:"author_mid"`
						AuthorName     string `json:"author_name"`
						PlaylistID     int64  `json:"playlist_id"`
						IsPlaylist     int    `json:"is_playlist"`
						GetPlaylistID  int64  `json:"get_playlist_id"`
						IsContribution int    `json:"is_contribution"`
						ExtraInfo      struct {
							Sceneid string `json:"sceneid"`
						} `json:"extra_info"`
						HasRecommendVideo int `json:"has_recommend_video"`
						BackPasterInfo    struct {
							HasBackPaster int `json:"has_back_paster"`
							RequestParam  struct {
								VideoType        int    `json:"video_type"`
								VideoOrientation string `json:"video_orientation"`
							} `json:"request_param"`
						} `json:"back_paster_info"`
						AuthorVerifiedType    int `json:"author_verified_type"`
						VideoDownloadStrategy struct {
							AbandonDownload int `json:"abandon_download"`
						} `json:"video_download_strategy"`
						Banner struct {
							URL       string `json:"url"`
							Scheme    string `json:"scheme"`
							Link      string `json:"link"`
							AppScheme string `json:"app_scheme"`
							Actionlog struct {
								Oid     string `json:"oid"`
								ActCode int    `json:"act_code"`
							} `json:"actionlog"`
						} `json:"banner"`
						OnlineUsers        string `json:"online_users"`
						OnlineUsersNumber  int    `json:"online_users_number"`
						TTL                int    `json:"ttl"`
						StorageType        string `json:"storage_type"`
						IsKeepCurrentMblog int    `json:"is_keep_current_mblog"`
					} `json:"media_info"`
					PlayCount int    `json:"play_count"`
					ObjectID  string `json:"object_id"`
				} `json:"page_info"`
				Pics []struct {
					Pid  string `json:"pid"`
					URL  string `json:"url"`
					Size string `json:"size"`
					Geo  struct {
						Width  int  `json:"width"`
						Height int  `json:"height"`
						Croped bool `json:"croped"`
					} `json:"geo"`
					Large struct {
						Size string `json:"size"`
						URL  string `json:"url"`
						Geo  struct {
							Width  string `json:"width"`
							Height string `json:"height"`
							Croped bool   `json:"croped"`
						} `json:"geo"`
					} `json:"large"`
				} `json:"pics"`
				Bid   string `json:"bid"`
				Title struct {
					Text      string `json:"text"`
					BaseColor int    `json:"base_color"`
				} `json:"title"`
			} `json:"mblog,omitempty"`
			ShowType int `json:"show_type"`
		} `json:"cards"`
		Scheme string `json:"scheme"`
	} `json:"data"`
}
type LongText struct {
	Ok   int `json:"ok"`
	Data struct {
		Ok              int    `json:"ok"`
		LongTextContent string `json:"longTextContent"`
		RepostsCount    int    `json:"reposts_count"`
		CommentsCount   int    `json:"comments_count"`
		AttitudesCount  int    `json:"attitudes_count"`
	} `json:"data"`
}

type weiboMsg struct {
	// -1 for net error
	// -2 for json error
	// 0 for old
	ok         int
	time       string
	text       string
	picNum     int
	pics       string
	video      bool
	videoPic   string
	videoUrl   string
	videoTitle string
	videoDesc  string
}

var newestId string = "nil"

func refresh() (wm weiboMsg) {
	//fmt.Println("entering refresh()")
	//defer fmt.Println("leaving refresh()")
	var re Weibo
	resp, err := http.Get("https://m.weibo.cn/api/container/getIndex?uid=5812573321&luicode=10000011&lfid=100103type%3D1%26q%3D%E5%B4%A9%E5%9D%8F3&type=uid&value=5812573321&containerid=1076035812573321")
	if err != nil {
		//fmt.Println(err)
		wm.ok = -1
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println(err)
		wm.ok = -1
		return
	}
	json.Unmarshal([]byte(string(body)), &re)

	// if not ok
	if re.Ok != 1 {
		wm.ok = -2
		return
	}

	// if newest
	if re.Data.Cards[1].Mblog.ID == newestId {
		wm.ok = 0
		return
	} else {
		newestId = re.Data.Cards[1].Mblog.ID
		wm.ok = 1
	}
	//fmt.Println(newestId)

	str := ""
	// if long text
	// testing long text 2 !!!
	if re.Data.Cards[1].Mblog.IsLongText {
		resp, err = http.Get("https://m.weibo.cn/statuses/extend?id=" + re.Data.Cards[1].Mblog.ID)
		body, err = ioutil.ReadAll(resp.Body)
		var lt LongText
		json.Unmarshal([]byte(string(body)), &lt)
		str = lt.Data.LongTextContent
	} else {
		str = re.Data.Cards[1].Mblog.Text
	}

	// edit string
	// <br /> -> \n
	for {
		pos1 := strings.Index(str, "<br")
		if pos1 == -1 {
			break
		}
		pos2 := pos1 + 6
		str = str[:pos1] + "\n" + str[pos2:]
	}
	// delete <...>
	for {
		pos1 := strings.Index(str, "<")
		pos2 := strings.Index(str, ">")
		if pos1 != -1 && pos2 != -1 {
			str = str[0:pos1] + str[pos2+1:]
		} else {
			//fmt.Println(str)
			break
		}
	}
	wm.text = str

	if strings.Contains(wm.text, "视频") {
		wm.video = true
		wm.videoPic = re.Data.Cards[1].Mblog.PageInfo.PagePic.URL
		wm.videoUrl = re.Data.Cards[1].Mblog.PageInfo.MediaInfo.Mp4HdURL
	}

	//editing time
	wm.time = re.Data.Cards[1].Mblog.CreatedAt

	// editing pics
	wm.picNum = re.Data.Cards[1].Mblog.PicNum
	for i := 0; i < wm.picNum; i++ {
		wm.pics += re.Data.Cards[1].Mblog.Pics[i].Large.URL + " "
	}

	return
}

func SendWeibo() {
	wm := refresh()
	//752390981
	data := Msg{"2325839514", 2, "547902826", "569927585"}
	if wm.ok != 1 {
		//SendMsg(data, strconv.Itoa(wm.ok))
		//fmt.Println(strconv.Itoa(wm.ok))
		return
	}
	str := wm.time + "\n" + wm.text

	picT := strings.Fields(wm.pics)
	for i := 0; i < wm.picNum; i++ {
		str += "\n" + "[IR:pic=" + picT[i] + "]"
	}

	SendMsg(data, str)
	if wm.video {
		content := "{\"config\":{\"forward\":true,\"type\":\"normal\",\"autosize\":true},\"prompt\":\"test\",\"app\":\"com.tencent.structmsg\",\"ver\":\"0.0.0.1\",\"view\":\"news\",\"meta\":{\"news\":{\"title\": \"" + wm.videoTitle + "\",\"desc\":\"" + wm.videoDesc + "\",\"preview\":\"" + wm.videoPic + "\",\"tag\":\"Athena\",\"jumpUrl\":\"" + wm.videoUrl + "\",\"appid\":1,\"app_type\":1,\"action\":\"\",\"source_url\":\"\",\"source_icon\":\"\",\"android_pkg_name\":\"com.logiase.top\"}},\"desc\":\"新闻\"}"
		SendJson(data, content)
	}
}

func WeiboTimer() {
	spec := "* */10 * * * ?"
	c := cron.New()
	c.AddFunc(spec, SendWeibo)
	c.Start()

	select {}
}
