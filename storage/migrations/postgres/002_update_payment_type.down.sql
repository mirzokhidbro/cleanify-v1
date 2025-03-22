-- Convert existing payment_type values back to text format
UPDATE transactions 
SET payment_type = CASE 
    WHEN payment_type::text = '1' THEN 'cach'
    WHEN payment_type::text = '2' THEN 'credit_card'
    ELSE 'cach'
END;

-- Change column type back to TEXT
ALTER TABLE transactions 
ALTER COLUMN payment_type TYPE TEXT;
