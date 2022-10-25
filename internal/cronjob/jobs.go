package cronjob

func RegisterJob(c *Cron, cronjob func()) {
	_, err := c.Cron.AddFunc("* * * * *", cronjob)
	if err != nil {
		c.Log.Error(err, "cron - registerJobList")
	}
}
