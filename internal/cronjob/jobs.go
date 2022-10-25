package cronjob

func RegisterJob(c *Cron, cronPattern string, cronjob func()) {
	_, err := c.Cron.AddFunc(cronPattern, cronjob)
	if err != nil {
		c.Log.Error(err, "cron - registerJobList")
	}
}
