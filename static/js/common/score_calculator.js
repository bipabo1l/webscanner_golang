/**
 * 积分计算器
 *
 * 自定义分值，不需要调用此计算器
 * 使用方式：
 *      参数1 分值
 *      参数2 等级
 *      ScoreClass.Calculation(10, 4)
 *
 *  @return int 0 计算错误，可能分值与等级不符合，{1:算法,2:应获积分}
 */
var ScoreClass = {
    //设置的分值，默认0
    useScore: 0,
    //分值区间
    scoreRange: {1: "1-3", 2: "4-6", 3: "7-8", 4: "9-10"},
    //系数
    ratioRange: {1: 9, 2: 15, 3: 60, 4: 120},
    //计算
    Calculation: function (score, level) {
        if (typeof(this.scoreRange[level]) != 'undefined') {
            this.useScore = score
            tmpScoreRange = this.scoreRange[level];
            err = this.checkScore(tmpScoreRange);
            if (err) {
                return {
                    levelL: this.ratioRange[level] + " * " + this.useScore,
                    scoreL: this.ratioRange[level] * this.useScore
                };
            } else {
                return 0;
            }
        } else {
            return 0;
        }
    },
    //校验等级与分值区间
    checkScore: function (tmpRange) {
        tmpRange = tmpRange.split("-");
        if (this.useScore >= parseInt(tmpRange[0]) && this.useScore <= parseInt(tmpRange[1])) {
            return true;
        } else {
            return false;
        }
    },
    //根据等级获取可用分值
    GetScore: function (level) {
        if (typeof(this.scoreRange[level]) != 'undefined') {
            tmpScoreRange = this.scoreRange[level];
            return tmpScoreRange;
        } else {
            return 0;
        }
    }
}
