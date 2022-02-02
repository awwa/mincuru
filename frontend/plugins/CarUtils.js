const makers = [
    {value: "トヨタ"},
    {value: "レクサス"},
    {value: "ホンダ"},
    {value: "日産"},
    {value: "マツダ"},
    {value: "スバル"},
    {value: "三菱"},
    {value: "スズキ"},
    {value: "ダイハツ"},
    {value: "フォルクスワーゲン"},
    {value: "アウディ"},
    {value: "BMW"},
    {value: "オペル"},
    {value: "ポルシェ"},
    {value: "メルセデス・ベンツ"},
    {value: "スマート"},
    {value: "アルピナ"},
    {value: "メルセデス-AMG"},
    {value: "ベントレー"},
    {value: "ジャガー"},
    {value: "ミニ"},
    {value: "ランドローバー"},
    {value: "フィアット"},
    {value: "アルファ・ロメオ"},
    {value: "ルノー"},
    {value: "シトロエン"},
    {value: "プジョー"},
    {value: "ボルボ"},
]

export default ({app}, inject) => {
    inject("makers", makers)
}